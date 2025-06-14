import React from 'react';
import './storage_usage.css';
import storageIcon from '../../assets/icons/storage.png';

const StorageUsage: React.FC = () => {
  return (
    <div className="health-card">
      <div className="health-card-header">
        <div className="health-icon-container health-gradient-bg">
          <img src={storageIcon} alt="Storage Icon" className="health-cpu-img" />
        </div>
        <div>
          <h4>Storage Usage</h4>
          <p>458GB / 1000GB</p>
        </div>
        <span className="health-expand">⤢</span>
      </div>
      <h2 className="health-value health-green">45.8% <span className="health-label">used</span></h2>
      <div className="health-bar-track">
        <div className="health-bar-fill" style={{ width: '45.8%' }}></div>
      </div>
      <div className="health-range">
        <span>Available: 542GB</span>
        <span>Used: 458GB</span>
      </div>
    </div>
  );
};

export default StorageUsage;