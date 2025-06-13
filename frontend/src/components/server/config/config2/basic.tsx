import React from 'react';
import './basic.css';

const Basic: React.FC = () => {
  return (
    <div className="card basic-card">
      <h3>Basic Information</h3>
      <ul>
        <li>Hostname: server-01</li>
        <li>IP Address: 192.168.0.101</li>
        <li>OS: Ubuntu 22.04</li>
      </ul>
    </div>
  );
};

export default Basic;
