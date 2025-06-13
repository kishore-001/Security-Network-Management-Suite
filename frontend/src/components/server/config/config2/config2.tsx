import React from 'react';
import './config2.css';
import Basic from './basic';
import FirewallManagement from './firewall_management';
import NetworkInterface from './network_interface';
import RouteTable from './route_table';

const Config2: React.FC = () => {
  return (
    <div className="config2-container">
      <h2 className="config2-title">Network Configuration</h2>
      <div className="config2-grid">
        <Basic />
        <FirewallManagement />
        <NetworkInterface />
        <RouteTable />
      </div>
    </div>
  );
};

export default Config2;
