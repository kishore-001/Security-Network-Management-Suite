import React from 'react';

interface Props {
  activeTab: 'user' | 'config';
  setActiveTab: (tab: 'user' | 'config') => void;
}

const SettingsTabs: React.FC<Props> = ({ activeTab, setActiveTab }) => {
  return (
 <div className="tab-wrapper">
  <div className="tab-container">
    <button
      className={`tab-btn ${activeTab === 'user' ? 'active' : ''}`}
      onClick={() => setActiveTab('user')}
    >
      User Management
    </button>
    <button
      className={`tab-btn ${activeTab === 'config' ? 'active' : ''}`}
      onClick={() => setActiveTab('config')}
    >
      Additional Configuration
    </button>
  </div>
</div>

  );
};

export default SettingsTabs;
