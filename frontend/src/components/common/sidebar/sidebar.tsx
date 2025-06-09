import React, { useState } from 'react';
import {
  FaCog,
  FaHeartbeat,
  FaDatabase,
  FaHdd,
  FaSlidersH
} from 'react-icons/fa';
import './sidebar.css';

const Sidebar: React.FC = () => {
  const [activeItem, setActiveItem] = useState('Settings');

  const menuItems = [
    { label: 'Configuration', icon: <FaSlidersH />, highlight: true },
    { label: 'Monitoring & Alerts', icon: <FaHeartbeat />, highlight: true },
    { label: 'Health Monitoring', icon: <FaHeartbeat /> },
    { label: 'Logging Systems', icon: <FaDatabase />, highlight: true },
    { label: 'Backup Management', icon: <FaHdd /> },
    { label: 'Settings', icon: <FaCog />, addSpacer: true, permanentPurple: true }
  ];

  return (
    <aside className="sidebar">
      <nav className="nav-menu">
        {menuItems.map(({ label, icon, highlight, addSpacer, permanentPurple }) => (
          <React.Fragment key={label}>
            {addSpacer && <div className="settings-spacer"></div>}
            <div
              className={`nav-item ${
                activeItem === label && !permanentPurple ? 'active' : ''
              } ${highlight ? 'highlight' : ''} ${
                permanentPurple ? 'settings-permanent' : ''
              }`}
              onClick={() => setActiveItem(label)}
            >
              <span className="nav-icon">{icon}</span>
              <span>{label}</span>
            </div>
          </React.Fragment>
        ))}
      </nav>
    </aside>
  );
};

export default Sidebar;