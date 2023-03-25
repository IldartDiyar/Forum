import React from 'react';
import ErrorIcon from '@mui/icons-material/Error';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';


export const PageNotFound = () => {
 
  return (
    <Paper elevation={3} style={{ padding: '2rem', margin: '2rem' }}>
      <Box display="flex" flexDirection="column" alignItems="center" justifyContent="center">    
        <ErrorIcon color="error" style={{ fontSize: '4rem' }} />
        
        <Typography variant="h3" fontFamily='Dosis' pb={5} color="error">
          404
        </Typography>
        
        <Typography variant="h4" fontFamily='Bebas Neue'>
          PAGE NOT FOUND
        </Typography>
        
      </Box>
    
    </Paper>
  )
}