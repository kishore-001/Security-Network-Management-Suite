import React from 'react';
import './issues.css';

const Issues: React.FC = () => {
  return (
    <div className="issues-card">
      <h3>Issues</h3>
      <div className="issues-stats">
        <div><span>11</span> Total</div>
        <div className="red"><span>10</span> New</div>
        <div className="yellow"><span>0</span> In Progress</div>
        <div className="green"><span>1</span> Resolved</div>
      </div>
      <div className="progress-bar">
        <div className="progress-fill" style={{ width: '9.1%' }}></div>
      </div>
      <div className="percentage">9.1%</div>
    </div>
  );
};

export default Issues;
