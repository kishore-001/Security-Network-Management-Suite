import React from 'react';
import './serveroverview.css';

const ServerOverview: React.FC = () => {
  return (
    <div className="server-overview">
      <div className="status-item">
        <span className="label">Server Status</span>
        <span className="value online">Online</span>
      </div>
      <div className="status-item">
        <span className="label">Uptime</span>
        <span className="value">19:00:00</span>
      </div>
      <div className="status-item">
        <span className="label">Active Users</span>
        <span className="value">3</span>
      </div>
    </div>
  );
};

export default ServerOverview;
