
import React from 'react';
import './firewall_management.css';

const FirewallManagement: React.FC = () => {
  return (
    <div className="card firewall-card">
      <h3>Firewall Rules</h3>
      <ul>
        <li>Allow: 22/tcp (SSH)</li>
        <li>Allow: 443/tcp (HTTPS)</li>
        <li>Deny: 23/tcp (Telnet)</li>
      </ul>
    </div>
  );
};

export default FirewallManagement;
