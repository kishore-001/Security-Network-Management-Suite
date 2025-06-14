import React from 'react';
import './ports_opened.css';
import globeIcon from '../../assets/icons/ports.png';

const PortsOpened: React.FC = () => {
  const ports = [80, 443, 920, 22, 3306, 5432];

  return (
    <div className="health-card">
      <div className="health-card-header">
        <div className="health-icon-container health-gradient-bg">
          <img src={globeIcon} alt="Ports Icon" className="health-cpu-img" />
        </div>
        <div>
          <h4>Ports Opened</h4>
          <p>8 active ports</p>
        </div>
        <span className="health-expand">⤢</span>
      </div>
      <div className="health-ports">
        {ports.map((port, index) => (
          <span key={index} className={`health-port-tag ${port <= 1024 ? 'health-port-critical' : ''}`}>{port}</span>
        ))}
        <p className="health-more-ports">+2 more ports</p>
      </div>
    </div>
  );
};

export default PortsOpened;