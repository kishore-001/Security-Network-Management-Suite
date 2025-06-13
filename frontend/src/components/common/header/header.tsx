import React, { useState, useRef, useEffect } from "react";
import "./header.css";
import { FaBell, FaNetworkWired, FaServer } from "react-icons/fa";
import { IoIosArrowDown } from "react-icons/io";
import { useNavigate, useLocation } from "react-router-dom";

const Header: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const [activeTab, setActiveTab] = useState<"network" | "server">("network");
  const [selectedServer, setSelectedServer] = useState("Server 1");
  const [dropdownOpen, setDropdownOpen] = useState(false);

  // Create a ref for the dropdown wrapper
  const dropdownRef = useRef<HTMLDivElement>(null);

  const servers = ["Server 1", "Server 2", "Server 3", "Server 4"];

  // Route to title mapping
  const getPageTitle = (pathname: string): string => {
    const routeTitles: { [key: string]: string } = {
      "/": "Configuration",
      "/login": "Login",
      "/health": "System Health",
      "/log": "Logging Systems",
      "/backup": "Backup Management",
      "/alert": "Monitoring & Alerts",
      "/resource": "Resource Optimization",
      "/settings": "Settings",
    };

    return routeTitles[pathname] || "SNSMS";
  };

  // Handle click outside to close dropdown
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node)
      ) {
        setDropdownOpen(false);
      }
    };

    // Add event listener when dropdown is open
    if (dropdownOpen) {
      document.addEventListener("mousedown", handleClickOutside);
    }

    // Cleanup event listener
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [dropdownOpen]);

  const toggleDropdown = () => setDropdownOpen((prev) => !prev);

  const handleServerSelect = (server: string) => {
    setSelectedServer(server);
    setDropdownOpen(false);
  };

  return (
    <div className="header-container">
      <div className="header-left">
        <div className="logo-circle">
          <span className="logo-letter">S</span>
          <span className="status-indicator" />
        </div>
        <div className="brand-info">
          <div className="brand-title">SNSMS</div>
          <div className="brand-subtitle">Network Management Suite</div>
        </div>

        {/* Dynamic Page Title */}
        <div className="page-title-section">
          <div className="page-title">{getPageTitle(location.pathname)}</div>
        </div>
      </div>

      <div className="header-right">
        {/* Add ref to the dropdown wrapper */}
        <div className="server-dropdown-wrapper" ref={dropdownRef}>
          <div className="server-dropdown" onClick={toggleDropdown}>
            {selectedServer}
            <IoIosArrowDown
              className={`dropdown-icon ${dropdownOpen ? "rotated" : ""}`}
            />
          </div>
          {dropdownOpen && (
            <div className="server-dropdown-menu">
              {servers.map((server) => (
                <div
                  key={server}
                  className="server-dropdown-item"
                  onClick={() => handleServerSelect(server)}
                >
                  {server}
                </div>
              ))}
            </div>
          )}
        </div>

        <div className="alert-icon" onClick={() => navigate("/alert")}>
          <FaBell className="bell-icon" />
          <span style={{ paddingLeft: "5px" }}>Alerts</span>
          <span className="alert-count">3</span>
        </div>

        <div className="toggle-switch">
          <div
            className={`toggle-option ${activeTab === "network" ? "active-network" : ""}`}
            onClick={() => setActiveTab("network")}
          >
            <FaNetworkWired />
            <span>Network</span>
          </div>
          <div
            className={`toggle-option ${activeTab === "server" ? "active-server" : ""}`}
            onClick={() => setActiveTab("server")}
          >
            <FaServer />
            <span>Server</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Header;
