import { useState } from "react";
import "./Setup.css";

/**
 * Setup guides the user through initial configuration when no user exists.
 * Multi-step wizard: Welcome -> Bazarr Import (optional) -> Admin User -> Server Settings -> Complete
 */
export default function Setup() {
  const [step, setStep] = useState(0);
  const [serverName, setServerName] = useState("Subtitle Manager");
  const [reverseProxy, setReverseProxy] = useState(false);
  const [adminUser, setAdminUser] = useState("");
  const [adminPass, setAdminPass] = useState("");

  // Bazarr import state
  const [bazarrURL, setBazarrURL] = useState("");
  const [bazarrAPIKey, setBazarrAPIKey] = useState("");
  const [bazarrSettings, setBazarrSettings] = useState(null);
  const [selectedSettings, setSelectedSettings] = useState({});
  const [bazarrLoading, setBazarrLoading] = useState(false);
  const [bazarrError, setBazarrError] = useState("");

  const [status, setStatus] = useState("");
  const [loading, setLoading] = useState(false);

  const importBazarr = async () => {
    if (!bazarrURL || !bazarrAPIKey) {
      setBazarrError("Please enter both URL and API key");
      return;
    }

    setBazarrLoading(true);
    setBazarrError("");

    try {
      const res = await fetch("/api/setup/bazarr", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          url: bazarrURL,
          api_key: bazarrAPIKey
        }),
      });

      if (res.ok) {
        const data = await res.json();
        setBazarrSettings(data.preview);
        // Pre-select all settings for import
        const selected = {};
        Object.keys(data.preview).forEach(key => {
          selected[key] = true;
        });
        setSelectedSettings(selected);
      } else {
        const errorText = await res.text();
        setBazarrError(errorText || "Failed to connect to Bazarr");
      }
    } catch (err) {
      setBazarrError("Network error: " + err.message);
    } finally {
      setBazarrLoading(false);
    }
  };

  const submit = async () => {
    setLoading(true);

    // Build the final configuration including Bazarr imports
    const integrations = {};
    const importedConfig = {};

    if (bazarrSettings) {
      Object.keys(selectedSettings).forEach(key => {
        if (selectedSettings[key]) {
          importedConfig[key] = bazarrSettings[key];
        }
      });
    }

    const body = {
      server_name: serverName,
      reverse_proxy: reverseProxy,
      admin_user: adminUser,
      admin_pass: adminPass,
      integrations,
      ...importedConfig // Merge Bazarr settings
    };

    try {
      const res = await fetch("/api/setup", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body),
      });

      if (res.ok) {
        setStatus("complete");
        setTimeout(() => window.location.reload(), 2000);
      } else {
        setStatus("error");
      }
    } catch (err) {
      setStatus("error");
    } finally {
      setLoading(false);
    }
  };

  const next = () => setStep(step + 1);
  const prev = () => setStep(step - 1);
  const skip = () => {
    setBazarrSettings(null);
    setSelectedSettings({});
    next();
  };

  if (status === "complete") {
    return (
      <div className="setup-container">
        <div className="setup-card">
          <div className="setup-success">
            <div className="success-icon">✅</div>
            <h1>Setup Complete!</h1>
            <p>Your Subtitle Manager is ready to use.</p>
            <p>Redirecting to login...</p>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="setup-container">
      <div className="setup-card">
        <div className="setup-header">
          <div className="setup-logo">
            <div className="subtitle-icon">
              <div className="subtitle-lines">
                <div className="line"></div>
                <div className="line"></div>
                <div className="line"></div>
              </div>
            </div>
          </div>
          <div className="setup-progress">
            <div className="progress-bar">
              <div
                className="progress-fill"
                style={{ width: `${(step + 1) * 25}%` }}
              ></div>
            </div>
            <span className="progress-text">Step {step + 1} of 4</span>
          </div>
        </div>

        <div className="setup-content">
          {step === 0 && (
            <div className="setup-step">
              <h1>Welcome to Subtitle Manager</h1>
              <p className="subtitle">
                A powerful, self-hosted subtitle management system that automates
                subtitle downloading, translation, and organization for your media library.
              </p>

              <div className="features-grid">
                <div className="feature">
                  <div className="feature-icon">🔄</div>
                  <h3>Automated Downloads</h3>
                  <p>Automatically find and download subtitles for your movies and TV shows</p>
                </div>
                <div className="feature">
                  <div className="feature-icon">🌐</div>
                  <h3>Translation</h3>
                  <p>Translate subtitles to any language using Google Translate or OpenAI</p>
                </div>
                <div className="feature">
                  <div className="feature-icon">🔌</div>
                  <h3>Integrations</h3>
                  <p>Works seamlessly with Sonarr, Radarr, Plex, and other media tools</p>
                </div>
                <div className="feature">
                  <div className="feature-icon">📱</div>
                  <h3>Modern UI</h3>
                  <p>Clean, responsive web interface accessible from any device</p>
                </div>
              </div>

              <div className="setup-actions">
                <button onClick={next} className="btn btn-primary">
                  Get Started
                </button>
              </div>
            </div>
          )}

          {step === 1 && (
            <div className="setup-step">
              <h1>Import from Bazarr</h1>
              <p className="subtitle">
                Already using Bazarr? Import your existing configuration to get started quickly.
              </p>

              <div className="bazarr-import">
                <div className="input-group">
                  <label htmlFor="bazarr-url">Bazarr URL</label>
                  <input
                    id="bazarr-url"
                    type="url"
                    placeholder="http://localhost:6767"
                    value={bazarrURL}
                    onChange={(e) => setBazarrURL(e.target.value)}
                    className="form-input"
                  />
                </div>

                <div className="input-group">
                  <label htmlFor="bazarr-key">API Key</label>
                  <input
                    id="bazarr-key"
                    type="password"
                    placeholder="Your Bazarr API key"
                    value={bazarrAPIKey}
                    onChange={(e) => setBazarrAPIKey(e.target.value)}
                    className="form-input"
                  />
                </div>

                {bazarrError && (
                  <div className="error-message">{bazarrError}</div>
                )}

                <button
                  onClick={importBazarr}
                  disabled={bazarrLoading}
                  className="btn btn-secondary"
                >
                  {bazarrLoading ? "Connecting..." : "Connect to Bazarr"}
                </button>

                {bazarrSettings && (
                  <div className="settings-preview">
                    <h3>Found Settings</h3>
                    <p>Select which settings to import:</p>
                    <div className="settings-list">
                      {Object.entries(bazarrSettings).map(([key, value]) => (
                        <label key={key} className="setting-item">
                          <input
                            type="checkbox"
                            checked={selectedSettings[key] || false}
                            onChange={(e) => setSelectedSettings({
                              ...selectedSettings,
                              [key]: e.target.checked
                            })}
                          />
                          <div className="setting-info">
                            <span className="setting-key">{key}</span>
                            <span className="setting-value">
                              {typeof value === 'object' ? JSON.stringify(value) : String(value)}
                            </span>
                          </div>
                        </label>
                      ))}
                    </div>
                  </div>
                )}
              </div>

              <div className="setup-actions">
                <button onClick={prev} className="btn btn-outline">Back</button>
                <button onClick={skip} className="btn btn-outline">Skip</button>
                <button
                  onClick={next}
                  disabled={!bazarrSettings}
                  className="btn btn-primary"
                >
                  Continue
                </button>
              </div>
            </div>
          )}

          {step === 2 && (
            <div className="setup-step">
              <h1>Create Admin Account</h1>
              <p className="subtitle">
                Create your administrator account to manage Subtitle Manager.
              </p>

              <div className="form-group">
                <div className="input-group">
                  <label htmlFor="admin-user">Username</label>
                  <input
                    id="admin-user"
                    type="text"
                    placeholder="admin"
                    value={adminUser}
                    onChange={(e) => setAdminUser(e.target.value)}
                    className="form-input"
                    required
                  />
                </div>

                <div className="input-group">
                  <label htmlFor="admin-pass">Password</label>
                  <input
                    id="admin-pass"
                    type="password"
                    placeholder="Choose a secure password"
                    value={adminPass}
                    onChange={(e) => setAdminPass(e.target.value)}
                    className="form-input"
                    required
                  />
                </div>
              </div>

              <div className="setup-actions">
                <button onClick={prev} className="btn btn-outline">Back</button>
                <button
                  onClick={next}
                  disabled={!adminUser || !adminPass}
                  className="btn btn-primary"
                >
                  Continue
                </button>
              </div>
            </div>
          )}

          {step === 3 && (
            <div className="setup-step">
              <h1>Server Configuration</h1>
              <p className="subtitle">
                Configure basic server settings for your deployment.
              </p>

              <div className="form-group">
                <div className="input-group">
                  <label htmlFor="server-name">Server Name</label>
                  <input
                    id="server-name"
                    type="text"
                    placeholder="Subtitle Manager"
                    value={serverName}
                    onChange={(e) => setServerName(e.target.value)}
                    className="form-input"
                  />
                </div>

                <div className="checkbox-group">
                  <label className="checkbox-label">
                    <input
                      type="checkbox"
                      checked={reverseProxy}
                      onChange={(e) => setReverseProxy(e.target.checked)}
                    />
                    <span className="checkmark"></span>
                    Running behind a reverse proxy
                  </label>
                  <p className="help-text">
                    Enable this if you're using nginx, Apache, or another reverse proxy
                  </p>
                </div>
              </div>

              {status === "error" && (
                <div className="error-message">
                  Setup failed. Please check your settings and try again.
                </div>
              )}

              <div className="setup-actions">
                <button onClick={prev} className="btn btn-outline">Back</button>
                <button
                  onClick={submit}
                  disabled={loading}
                  className="btn btn-primary"
                >
                  {loading ? "Setting up..." : "Complete Setup"}
                </button>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
