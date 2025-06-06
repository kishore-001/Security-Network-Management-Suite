import React from 'react';
import "./App.css"
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
// Update the import path below if Sidebar is located elsewhere, e.g. './components/Sidebar'
// Update the import path below if Sidebar is located elsewhere, e.g. './Sidebar' or './pages/Sidebar'
// Update the path below to the correct location and filename of Sidebar
import Header from './components/common/header/header';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        {/* Redirect root to /sidebar */}
        <Route path="/" element={<Navigate to="/header  " replace />} />
        
        {/* Sidebar route */}
        <Route path="/header" element={<Header />} />
      </Routes>
    </Router>
  );
};

export default App