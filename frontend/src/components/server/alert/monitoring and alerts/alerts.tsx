import React from 'react';
import './alerts.css';

const Alerts: React.FC = () => {
  return (
    <div className="alerts-card">
      <h3>Alerts</h3>
      <div className="circle-alert">
        <div className="circle">1551</div>
        <span><strong>Total</strong></span>
        <span className="alert-counts">
          <div style={{ fontSize: "20px" }} className="red_text"><strong>1401 High</strong></div>
          <div style={{ fontSize: "20px" }} className="yellow"><strong>39 Medium</strong></div>
          <div  style={{ fontSize: "20px" }} className="blue"><strong>111 Low</strong></div>
        </span>
      </div>
    </div>
  );
};

export default Alerts;
