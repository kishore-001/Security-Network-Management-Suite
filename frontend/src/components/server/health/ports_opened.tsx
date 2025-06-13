import React from 'react';
import './ports_opened.css';

const ports = [80, 443, 920, 22, 3306, 5432];

const PortsOpened: React.FC = () => {
  return (
    <div className="card ports-card">
      <div className="card-header">
        <div className="icon ports-icon" />
        <div>
          <h4>Ports Opened</h4>
          <p>{ports.length} active ports</p>
        </div>
        <span className="expand">⤢</span>
      </div>
      <div className="card-body ports-body">
        <div className="ports-grid">
          {ports.map((port, idx) => (
            <div key={idx} className={`port-box ${port === 80 || port === 443 || port === 22 ? 'critical' : ''}`}>
              {port}
            </div>
          ))}
        </div>
        <p className="more-ports">+2 more ports</p>
      </div>
    </div>
  );
};

export default PortsOpened;
