import { useState } from "react";

/**
 * Convert provides a form to upload a subtitle file which is
 * converted to SRT format via the /api/convert endpoint.
 * The resulting file is downloaded by the browser.
 */
export default function Convert() {
  const [file, setFile] = useState(null);
  const [status, setStatus] = useState("");

  const doConvert = async () => {
    if (!file) return;
    const form = new FormData();
    form.append("file", file);
    const res = await fetch("/api/convert", { method: "POST", body: form });
    if (res.ok) {
      const blob = await res.blob();
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement("a");
      a.href = url;
      a.download = "converted.srt";
      a.click();
      setStatus("converted");
    } else {
      setStatus("error");
    }
  };

  return (
    <div className="convert">
      <h1>Convert Subtitle</h1>
      <input data-testid="file" type="file" onChange={(e) => setFile(e.target.files[0])} />
      <button onClick={doConvert}>Convert</button>
      <p>{status}</p>
    </div>
  );
}
