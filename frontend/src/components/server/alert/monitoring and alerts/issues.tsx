import React from 'react';
import './issues.css';

const Issues: React.FC = () => {
  return (
    <div className="issues-card">
      <h3>Issues</h3>
      <div className="issues-stats">
        <div style={{ fontSize: "20px" }}><span>11</span> Total</div>
        <div style={{ fontSize: "20px" }} className="red_text"><span>10</span> New</div>
        <div style={{ fontSize: "20px" }}className="yellow"><span>0</span> In Progress</div>
        <div style={{ fontSize: "20px" }} className="green"><span>1</span> Resolved</div>
      </div>
      <div className="progress-bar">
        <div className="progress-fill" style={{ width: '9.1%' }}></div>
      </div>
      <div className="percentage">9.1%</div>
    </div>
  );
};

export default Issues;
