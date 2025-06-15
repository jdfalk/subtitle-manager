// file: webui/src/UserManagement.jsx
import { useEffect, useState } from "react";
import {
  Box,
  Button,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";

/**
 * UserManagement displays all users and allows password resets.
 */
export default function UserManagement() {
  const [users, setUsers] = useState([]);

  const loadUsers = async () => {
    const res = await fetch("/api/users");
    if (res.ok) setUsers(await res.json());
  };

  const reset = async (id) => {
    if (!window.confirm("Reset password for this user?")) return;
    const res = await fetch(`/api/users/${id}/reset`, { method: "POST" });
    if (res.ok) {
      alert("Password reset and emailed");
    }
  };

  useEffect(() => {
    loadUsers();
  }, []);

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        Users
      </Typography>
      <Table size="small">
        <TableHead>
          <TableRow>
            <TableCell>Username</TableCell>
            <TableCell>Email</TableCell>
            <TableCell>Role</TableCell>
            <TableCell />
          </TableRow>
        </TableHead>
        <TableBody>
          {users.map((u) => (
            <TableRow key={u.id}>
              <TableCell>{u.username}</TableCell>
              <TableCell>{u.email}</TableCell>
              <TableCell>{u.role}</TableCell>
              <TableCell>
                <Button size="small" onClick={() => reset(u.id)}>
                  Reset Password
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </Box>
  );
}
