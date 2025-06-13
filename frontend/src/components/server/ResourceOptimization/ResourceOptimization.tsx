import React from 'react';
import UsageStats from './UsageStats';
import ResourceCard from './ResourceCard';
import ServiceGrid from './ServiceGrid';
import './resource.css';

const ResourceOptimization: React.FC = () => {
  return (
    <div className="resource-dashboard">
      <div className="header">
        <div>
          <h2>Resource Optimization</h2>
          <p>System performance monitoring and optimization tools</p>
        </div>
        <div className="header-actions">
          <button className="btn">Refresh Data</button>
          <button className="btn">Export Report</button>
        </div>
      </div>

      <div className="metrics-section">
        <UsageStats label="Memory usage" value={39} />
        <UsageStats label="CPU usage" value={6.27} />
        <UsageStats label="Disk usage" value={6.89} />
        <div className="uptime">
          <div className="uptime-clock">🕒</div>
          <p className="uptime-value">00:02:04:34</p>
          <p className="uptime-label">Server Uptime</p>
        </div>
      </div>

      <div className="resource-cards">
        <ResourceCard
          icon="🗑️"
          title="Temporary File Cleanup"
          description="Clean system temporary files and cache"
          stats={["2.3 GB", "3 days ago"]}
        />
        <ResourceCard
          icon="🔄"
          title="Restart Services"
          description="Restart system services for optimization"
          stats={["10", "2 hours ago"]}
        />
      </div>

      <div className="service-overview">
        <h3>Service Architecture Overview</h3>
        <ServiceGrid />
      </div>
    </div>
  );
};

export default ResourceOptimization;
