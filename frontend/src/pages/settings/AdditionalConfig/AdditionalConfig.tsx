import React from 'react';
import MACListManager from './MacListManager';
import ConfigManagement from './ConfigManagement';
import './AdditionalConfig.css';

const AdditionalConfig: React.FC = () => {
  return (
    <div className="additional-config-container">
      <MACListManager />
      <ConfigManagement />
    </div>
  );
};

export default AdditionalConfig;