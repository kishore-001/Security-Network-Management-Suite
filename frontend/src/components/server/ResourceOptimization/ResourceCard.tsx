import React from 'react';
import './resourceCard.css';

interface ResourceCardProps {
  icon: string;
  title: string;
  description: string;
  stats: string[];
}

const ResourceCard: React.FC<ResourceCardProps> = ({ icon, title, description, stats }) => {
  return (
    <div className="resource-card">
      <div className="resource-icon">{icon}</div>
      <div className="resource-info">
        <h4>{title}</h4>
        <p>{description}</p>
        <div className="resource-stats">
          <span>{stats[0]}</span>
          <span>{stats[1]}</span>
        </div>
      </div>
    </div>
  );
};

export default ResourceCard;
