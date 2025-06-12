import { useState } from "react";

/**
 * Translate provides a form to upload a subtitle file and request
 * translation to a target language via the /api/translate endpoint.
 * The translated file is downloaded by the browser.
 */
export default function Translate() {
  const [file, setFile] = useState(null);
  const [lang, setLang] = useState("es");
  const [status, setStatus] = useState("");

  const doTranslate = async () => {
    if (!file) return;
    const form = new FormData();
    form.append("file", file);
    form.append("lang", lang);
    const res = await fetch("/api/translate", { method: "POST", body: form });
    if (res.ok) {
      const blob = await res.blob();
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement("a");
      a.href = url;
      a.download = "translated.srt";
      a.click();
      setStatus("translated");
    } else {
      setStatus("error");
    }
  };

  return (
    <div className="translate">
      <h1>Translate Subtitle</h1>
      <input data-testid="file" type="file" onChange={(e) => setFile(e.target.files[0])} />
      <input
        placeholder="Language"
        value={lang}
        onChange={(e) => setLang(e.target.value)}
      />
      <button onClick={doTranslate}>Translate</button>
      <p>{status}</p>
    </div>
  );
}
