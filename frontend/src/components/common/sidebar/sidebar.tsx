import { useNavigate, useLocation } from "react-router-dom";
import "./sidebar.css";
import icons from "../../../assets/icons";

interface Metric {
  icon: string;
  value: string;
  label: string;
  change: string;
  changeType: "positive" | "negative";
}

const metrics: Metric[] = [
  {
    icon: "📈",
    value: "1.2 Gbps",
    label: "Total Bandwidth",
    change: "+12%",
    changeType: "positive",
  },
  {
    icon: "🌐",
    value: "446",
    label: "Active Connections",
    change: "+5%",
    changeType: "positive",
  },
  {
    icon: "🛡️",
    value: "23",
    label: "Security Events",
    change: "-8%",
    changeType: "negative",
  },
  {
    icon: "💓",
    value: "98.5%",
    label: "System Health",
    change: "+2%",
    changeType: "positive",
  },
];

interface MenuItem {
  label: string;
  icon: keyof typeof icons;
  path: string;
  count?: number;
  alert?: boolean;
}

const menuItems: MenuItem[] = [
  { label: "Configuration", icon: "config", path: "/", count: 12 },
  {
    label: "Monitoring & Alerts",
    icon: "reports",
    path: "/alert",
    count: 3,
    alert: true,
  },
  {
    label: "Resource Optimization",
    icon: "reports",
    path: "/resource",
    count: 8,
  },
  { label: "Logging Systems", icon: "logg", path: "/log", count: 156 },
  { label: "Backup Management", icon: "backup", path: "/backup", count: 5 },
  { label: "Health", icon: "health", path: "/health" },
];

const Sidebar = () => {
  const navigate = useNavigate();
  const location = useLocation();

  return (
    <div className="sidebar-container">
      <div className="metrics-section">
        {metrics.map((metric, index) => (
          <div key={index} className={`metric-card metric-${index}`}>
            <div className="metric-change" data-type={metric.changeType}>
              {metric.change}
            </div>
            <div className="metric-value">{metric.value}</div>
            <div className="metric-label">{metric.label}</div>
          </div>
        ))}
      </div>

      <div className="menu-section">
        {menuItems.map((item, index) => {
          const isActive = location.pathname === item.path;

          return (
            <div
              key={index}
              className={`menu-item ${isActive ? "active-blue" : ""}`}
              onClick={() => navigate(item.path)}
            >
              <img
                src={icons[item.icon]}
                alt={item.label}
                className="menu-icon"
              />
              <span className="menu-label">{item.label}</span>
              {item.count !== undefined && (
                <span className={`menu-count ${item.alert ? "alert" : ""}`}>
                  {item.count}
                </span>
              )}
            </div>
          );
        })}

        {/* Settings */}
        <div
          className={`menu-item settings-button ${location.pathname === "/settings" ? "active-blue" : ""}`}
          onClick={() => navigate("/settings")}
        >
          <img src={icons.settings} alt="Settings" className="menu-icon" />
          <span className="menu-label">Settings</span>
        </div>
      </div>
    </div>
  );
};

export default Sidebar;
