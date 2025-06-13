import React from 'react';
import './network_interface.css';

const NetworkInterface: React.FC = () => {
  return (
    <div className="card network-card">
      <h3>Network Interfaces</h3>
      <ul>
        <li>eth0: 192.168.0.101</li>
        <li>wlan0: Not Connected</li>
      </ul>
    </div>
  );
};

export default NetworkInterface;
