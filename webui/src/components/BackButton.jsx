// file: webui/src/components/BackButton.jsx
// version: 1.0.0
// guid: 1d2f3e4a-5b6c-7d8e-9f01-23456789abcd

import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import Button from '@mui/material/Button';
import { useNavigate } from 'react-router-dom';

/**
 * BackButton navigates to the previous page or home if history is empty.
 * @returns {JSX.Element} Back navigation button
 */
export default function BackButton() {
  const navigate = useNavigate();

  const handleClick = () => {
    if (window.history.length > 1) {
      navigate(-1);
    } else {
      navigate('/');
    }
  };

  return (
    <Button onClick={handleClick} startIcon={<ArrowBackIcon />} sx={{ mb: 2 }}>
      Back
    </Button>
  );
}
