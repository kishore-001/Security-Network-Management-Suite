import React from 'react';
import './serveroverview.css';

const ServerOverview: React.FC = () => {
  return (
    <div className="server-overview">
      <h2 className="overview-title">Server Overview</h2>
      <div className="overview-metrics">
        <div className="metric-block">
          <div className="metric-value online">Online</div>
          <div className="metric-label">Server Status</div>
        </div>
        <div className="metric-block">
          <div className="metric-value uptime">19:00:00</div>
          <div className="metric-label">Uptime</div>
        </div>
        <div className="metric-block">
          <div className="metric-value users">3</div>
          <div className="metric-label">Active Users</div>
        </div>
      </div>
    </div>
  );
};

export default ServerOverview;
