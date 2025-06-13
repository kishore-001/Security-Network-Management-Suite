import React from 'react';
import './health.css';
import CpuUsage from './cpu_usage';
import MemoryUsage from './memory_usage';
import StorageUsage from './storage_usage';
import NetworkTraffic from './network_traffic';
import PortsOpened from './ports_opened';

const Health: React.FC = () => {
  return (
    <div className="health-container">
      <div className="health-header">
        <div>
          <h2 className="dashboard-title">System Health Dashboard</h2>
          <p className="dashboard-subtitle">Real-time monitoring of system resources and performance metrics</p>
        </div>
        <div className="header-buttons">
          <button className="refresh-btn">🔄 Refresh Data</button>
        </div>
      </div>

      <div className="top-row">
        <CpuUsage />
        <MemoryUsage />
        <StorageUsage />
      </div>

      <div className="bottom-row">
        <NetworkTraffic />
        <PortsOpened />
      </div>
    </div>
  );
};

export default Health;
