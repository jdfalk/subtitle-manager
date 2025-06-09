import { useEffect, useState } from "react";

/**
 * Settings renders a simple configuration form. Values are loaded from
 * `/api/config` and submitted back via POST to the same endpoint.
 */
export default function Settings() {
  const [config, setConfig] = useState(null);
  const [status, setStatus] = useState("");

  useEffect(() => {
    fetch("/api/config")
      .then((r) => r.json())
      .then(setConfig);
  }, []);

  const handleChange = (key, value) => {
    setConfig({ ...config, [key]: value });
  };

  const save = async () => {
    const res = await fetch("/api/config", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(config),
    });
    setStatus(res.ok ? "saved" : "error");
  };

  if (!config) return <p>Loading...</p>;

  return (
    <div className="settings">
      <h1>Settings</h1>
      <table>
        <tbody>
          {Object.entries(config).map(([k, v]) => (
            <tr key={k}>
              <td>{k}</td>
              <td>
                <input
                  value={v}
                  onChange={(e) => handleChange(k, e.target.value)}
                />
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      <button onClick={save}>Save</button>
      <p>{status}</p>
    </div>
  );
}
