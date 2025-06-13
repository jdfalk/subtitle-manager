import { useState } from "react";

/**
 * Setup guides the user through initial configuration when no user exists.
 * It captures the server name, admin credentials and basic integrations.
 */
export default function Setup() {
  const [step, setStep] = useState(0);
  const [serverName, setServerName] = useState("Subtitle Manager");
  const [reverseProxy, setReverseProxy] = useState(false);
  const [adminUser, setAdminUser] = useState("");
  const [adminPass, setAdminPass] = useState("");
  const [integrations, setIntegrations] = useState({
    sonarr: false,
    radarr: false,
    bazarr: false,
    plex: false,
    notifications: false,
  });
  const [status, setStatus] = useState("");

  const submit = async () => {
    const body = {
      server_name: serverName,
      reverse_proxy: reverseProxy,
      admin_user: adminUser,
      admin_pass: adminPass,
      integrations,
    };
    const res = await fetch("/api/setup", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body),
    });
    if (res.ok) {
      setStatus("setup complete");
    } else {
      setStatus("error");
    }
  };

  const next = () => setStep(step + 1);
  const prev = () => setStep(step - 1);

  return (
    <div className="setup">
      <h1>Initial Setup</h1>
      <p>
        Already using Bazarr? Run
        <code>subtitle-manager import-bazarr http://localhost:6767 MY_API_KEY</code>
        to copy your settings.
      </p>
      {step === 0 && (
        <div>
          <p>Server settings</p>
          <input
            placeholder="Server Name"
            value={serverName}
            onChange={(e) => setServerName(e.target.value)}
          />
          <label>
            <input
              type="checkbox"
              checked={reverseProxy}
              onChange={(e) => setReverseProxy(e.target.checked)}
            />
            Behind reverse proxy
          </label>
          <button onClick={next}>Next</button>
        </div>
      )}
      {step === 1 && (
        <div>
          <p>Create admin user</p>
          <input
            placeholder="Username"
            value={adminUser}
            onChange={(e) => setAdminUser(e.target.value)}
          />
          <input
            type="password"
            placeholder="Password"
            value={adminPass}
            onChange={(e) => setAdminPass(e.target.value)}
          />
          <button onClick={prev}>Back</button>
          <button onClick={next}>Next</button>
        </div>
      )}
      {step === 2 && (
        <div>
          <p>Enable integrations</p>
          {Object.keys(integrations).map((k) => (
            <label key={k} style={{ display: "block" }}>
              <input
                type="checkbox"
                checked={integrations[k]}
                onChange={(e) =>
                  setIntegrations({ ...integrations, [k]: e.target.checked })
                }
              />
              {k}
            </label>
          ))}
          <button onClick={prev}>Back</button>
          <button onClick={submit}>Finish</button>
        </div>
      )}
      <p>{status}</p>
    </div>
  );
}
