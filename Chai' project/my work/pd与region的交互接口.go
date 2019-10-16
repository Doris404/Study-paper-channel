// handleRegionHeartbeat updates the region information.
func (c *clusterInfo) handleRegionHeartbeat(region *core.RegionInfo) error {
	region = region.Clone()
	c.RLock()
	origin := c.Regions.GetRegion(region.GetId())
	isWriteUpdate, writeItem := c.CheckWriteStatus(region)
	isReadUpdate, readItem := c.CheckReadStatus(region)
	c.RUnlock()

	// Save to KV if meta is updated.
	// Save to cache if meta or leader is updated, or contains any down/pending peer.
	// Mark isNew if the region in cache does not have leader.
	var saveKV, saveCache, isNew bool
	if origin == nil {
		log.Infof("[region %d] Insert new region {%v}", region.GetId(), region)
		saveKV, saveCache, isNew = true, true, true
	} else {
		r := region.GetRegionEpoch()
		o := origin.GetRegionEpoch()
		// Region meta is stale, return an error.
		if r.GetVersion() < o.GetVersion() || r.GetConfVer() < o.GetConfVer() {
			return errors.Trace(ErrRegionIsStale(region.Region, origin.Region))
		}
		if r.GetVersion() > o.GetVersion() {
			log.Infof("[region %d] %s, Version changed from {%d} to {%d}", region.GetId(), core.DiffRegionKeyInfo(origin, region), o.GetVersion(), r.GetVersion())
			saveKV, saveCache = true, true
		}
		if r.GetConfVer() > o.GetConfVer() {
			log.Infof("[region %d] %s, ConfVer changed from {%d} to {%d}", region.GetId(), core.DiffRegionPeersInfo(origin, region), o.GetConfVer(), r.GetConfVer())
			saveKV, saveCache = true, true
		}
		if region.Leader.GetId() != origin.Leader.GetId() {// leader发生转移
			log.Infof("[region %d] Leader changed from {%v} to {%v}", region.GetId(), origin.GetPeer(origin.Leader.GetId()), region.GetPeer(region.Leader.GetId()))
			if origin.Leader.GetId() == 0 {
				isNew = true
			}
			saveCache = true
		}
		if len(region.DownPeers) > 0 || len(region.PendingPeers) > 0 {
			saveCache = true
		}
		if len(origin.DownPeers) > 0 || len(origin.PendingPeers) > 0 {
			saveCache = true
		}
		if region.ApproximateSize != origin.ApproximateSize {
			saveCache = true
		}
	}

	if saveKV && c.kv != nil {
		if err := c.kv.SaveRegion(region.Region); err != nil {
			// Not successfully saved to kv is not fatal, it only leads to longer warm-up
			// after restart. Here we only log the error then go on updating cache.
			log.Errorf("[region %d] fail to save region %v: %v", region.GetId(), region, err)
		}
	}
	if !isWriteUpdate && !isReadUpdate && !saveCache && !isNew {
		return nil
	}

	c.Lock()
	defer c.Unlock()
	if isNew {
		c.activeRegions++
	}

	if saveCache {
		overlaps := c.Regions.SetRegion(region)
		if c.kv != nil {
			for _, item := range overlaps {
				if err := c.kv.DeleteRegion(item); err != nil {
					log.Errorf("[region %d] fail to delete region %v: %v", item.GetId(), item, err)
				}
			}
		}
		for _, item := range overlaps {
			if c.regionStats != nil {
				c.regionStats.clearDefunctRegion(item.GetId())
			}
			c.labelLevelStats.clearDefunctRegion(item.GetId())
		}

		// Update related stores.
		if origin != nil {
			for _, p := range origin.Peers {
				c.updateStoreStatus(p.GetStoreId())
			}
		}
		for _, p := range region.Peers {
			c.updateStoreStatus(p.GetStoreId())
		}
	}

	if c.regionStats != nil {
		c.regionStats.Observe(region, c.getRegionStores(region))
	}

	key := region.GetId()
	if isWriteUpdate {
		c.HotCache.Update(key, writeItem, schedule.WriteFlow)
	}
	if isReadUpdate {
		c.HotCache.Update(key, readItem, schedule.ReadFlow)
	}
	return nil
}