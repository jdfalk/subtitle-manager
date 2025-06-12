import { useEffect, useState } from "react";

/**
 * History displays translation and download history with optional language filtering.
 * Records are loaded from `/api/history` and filtered client side.
 */
export default function History() {
  const [data, setData] = useState({ translations: [], downloads: [] });
  const [lang, setLang] = useState("");

  useEffect(() => {
    fetch("/api/history")
      .then((r) => r.json())
      .then(setData);
  }, []);

  const translations = data.translations.filter(
    (r) => !lang || r.Language === lang,
  );
  const downloads = data.downloads.filter((r) => !lang || r.Language === lang);

  return (
    <div className="history">
      <h1>History</h1>
      <input
        placeholder="Filter language"
        value={lang}
        onChange={(e) => setLang(e.target.value)}
      />
      <h2>Translations</h2>
      <table>
        <thead>
          <tr>
            <th>File</th>
            <th>Language</th>
            <th>Service</th>
          </tr>
        </thead>
        <tbody>
          {translations.map((t) => (
            <tr key={t.ID}>
              <td>{t.File}</td>
              <td>{t.Language}</td>
              <td>{t.Service}</td>
            </tr>
          ))}
        </tbody>
      </table>
      <h2>Downloads</h2>
      <table>
        <thead>
          <tr>
            <th>Video</th>
            <th>Language</th>
            <th>Provider</th>
          </tr>
        </thead>
        <tbody>
          {downloads.map((d) => (
            <tr key={d.ID}>
              <td>{d.VideoFile}</td>
              <td>{d.Language}</td>
              <td>{d.Provider}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
