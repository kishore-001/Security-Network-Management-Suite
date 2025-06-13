import React from 'react';
import './storage_usage.css';
import box from '../../../assets/icon/box.svg';
const StorageUsage: React.FC = () => {
  return (
    <div className="card storage-card">
      <div className="card-header">
         <div className="icon-container gradient_storage">
          <img src={box} alt="box Icon" className="img" />
        </div>
        <div>
          <h4>Storage Usage</h4>
          <p>458GB / 1000GB</p>
        </div>
        <span className="expand">⤢</span>
      </div>
      <div className="card-body">
        <h2 className="value green">45.8% <span className="label">used</span></h2>
        <div className="bar-track"><div className="bar-fill" style={{ width: '45.8%' }}></div></div>
        <div className="range">
          <span>Available: 542GB</span>
          <span>Used: 458GB</span>
        </div>
      </div>
    </div>
  );
};

export default StorageUsage;
