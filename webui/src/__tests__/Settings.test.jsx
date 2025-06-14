// file: webui/src/__tests__/Settings.test.jsx
import "@testing-library/jest-dom/vitest";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { beforeEach, describe, expect, test, vi } from "vitest";
import Settings from "../Settings.jsx";

describe("Settings component", () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn(() => Promise.resolve({ ok: true, json: () => Promise.resolve({}) }));
  });

  test("edits general settings", async () => {
    fetch
      .mockResolvedValueOnce({
        ok: true,
        json: () => Promise.resolve({ server_name: "Test" }),
      })
      .mockResolvedValueOnce({ ok: true, json: () => Promise.resolve([]) });

    render(<Settings />);

    fireEvent.click(screen.getByRole("tab", { name: /General/i }));
    await screen.findByDisplayValue("Test");

    fetch.mockResolvedValueOnce({ ok: true });
    fireEvent.change(screen.getByLabelText("Server Name"), {
      target: { value: "New" },
    });
    fireEvent.click(screen.getByText("Save"));

    await waitFor(() =>
      expect(fetch).toHaveBeenLastCalledWith(
        "/api/config",
        expect.objectContaining({ method: "POST" }),
      ),
    );
  });
});
