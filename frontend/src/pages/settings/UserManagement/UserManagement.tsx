import React from 'react';
import UserTable from './Usertable';
import './UserManagement.css';

const UserManagement: React.FC = () => {
  return (
    <div>
      <div className="action-buttons">
        <button className="btn add">+ Add User</button>
        <button className="btn remove">Remove User</button>
      </div>
      <UserTable />
    </div>
  );
};

export default UserManagement;
