import React from 'react';
import './health.css';
import CpuUsage from './cpu_usage';
import MemoryUsage from './memory_usage';
import StorageUsage from './storage_usage';
import NetworkTraffic from './network_traffic';
import PortsOpened from './ports_opened';

const Health: React.FC = () => {
  return (
    <div className="health-dashboard">
      <div className="health-header">
        <div>
          <h2>System Health Dashboard</h2>
          <p>Real-time monitoring of system resources and performance metrics</p>
        </div>
        <button className="health-refresh">Refresh Data</button>
      </div>
      <div className="health-grid">
        <CpuUsage />
        <MemoryUsage />
        <StorageUsage />
        <NetworkTraffic />
        <PortsOpened />
      </div>
    </div>
  );
};

export default Health;
