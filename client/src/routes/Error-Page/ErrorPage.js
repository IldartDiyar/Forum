import React from 'react';
import { useRouteError } from 'react-router';
import ErrorIcon from '@mui/icons-material/Error';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import { useNavigate } from 'react-router';
import { Button } from '@mui/material';

export const ErrorPage = () => {
  const error = useRouteError();
  const navigate = useNavigate();

  const navigateToHome = () => {
    navigate('/');
  };      
    return (
      <Paper elevation={3} style={{ padding: '2rem', margin: '2rem' }}>
        <Box display="flex" flexDirection="column" alignItems="center">
          
          <ErrorIcon color="error" style={{ fontSize: '4rem' }} />
          
          <Typography variant="h3" fontFamily='Dosis' pb={5} color="error">
            Error {error.status? error.status: ""}
          </Typography>
          
          <Typography variant="h6">
            {error.message ? error.message: ''}          
          </Typography>
        
          <Button
            variant="contained"
            color="primary"
            onClick={navigateToHome}
            sx={{ marginTop: '2rem' }}
          >
            HOME
          </Button>
        </Box>
      </Paper>
    )
}