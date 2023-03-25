import React, { useState } from 'react';
import { Alert, Avatar, Button, Input, Typography } from '@mui/material';
import {Grid, Box} from '@mui/material';
import { useLoaderData, useNavigate } from 'react-router';
import { Divider } from '@mui/material';
import { AddAPhoto } from '@mui/icons-material';
import { uploadImage } from '../../api/Profile/ProfileAPI';

const ProfilePage = () => {
  const {data} = useLoaderData();
  const [image, setImage] = useState(null);
  const fileSizeLimit = 5000000; //5 MB
  const [status, setStatus] = useState('');
  const [error, setError] = useState('');
  const [imageUrl, setImageUrl] = useState(null);
  const navigate = useNavigate();
  
  const handleImageChange = (e) => {
    const uploadedFile = e.target.files[0];
    setStatus('')
    if (uploadedFile && uploadedFile.size > fileSizeLimit){
      setError(`File size must be less than ${fileSizeLimit} bytes`)
      setImage(null);
    } else{
      setError('')
      setImage(e.target.files[0]);
      setImageUrl(URL.createObjectURL(e.target.files[0]));
    }
  };

  return (
      <Grid 
      container
      direction="row"
      justifyContent="center"
      alignItems="center"
      >
          <Grid item xs={12} > {/*HEADER OF A PAGE*/} 
            <Typography variant="h4" fontFamily='Dosis' sx={{display:'flex', justifyContent:'center', alignItems:'center', pt: 4, pb: 4}}>PROFILE PAGE</Typography>
            <Divider>
              <Typography variant="h4" fontFamily='Bebas Neue' sx={{display:'flex', justifyContent:'center', alignItems:'center', pt:2}}>
                  {data.Username && data.Username}
              </Typography>
            </Divider>  
          </Grid>
          {/* AVATAR */}
          <Grid item xs={6} sx={{height:'30vh', display:'flex', justifyContent:'center', alignItems:'center'}}>
            <Avatar alt="Profile image" src={data.Avatar!=="null" ? data.Avatar : null} sx={{ width:{ xs:100}, height: {xs:100}}}/>  
          </Grid>
          {/* PERSONAL INFO */}
          <Grid item xs={6} sx={{height:'30vh', display:'flex', flexDirection:'column', justifyContent:'space-evenly', alignItems:'center', position:'relative'}}>
            <Typography variant="h5" fontFamily='Dosis'>Username: {data.Username && data.Username}</Typography>
            <Typography variant="h5" fontFamily='Dosis'>Email: {data.Email && data.Email}</Typography>
            <Typography variant="h5" fontFamily='Dosis'>Posts: {data.MyPosts && data.MyPosts}</Typography>
            <Typography variant="h5" fontFamily='Dosis'>Received Likes: {data.LikesOnMyPost && data.LikesOnMyPost}</Typography>
            <Typography variant="h5" fontFamily='Dosis'>Liked Posts: {data.LikedPosts && data.LikedPosts}</Typography>
          </Grid>
          {/* IMAGE UPLOAD */}
          <Grid item xs={12}>
            <Divider />
            <Typography variant="h4" fontFamily='Dosis' sx={{display:'flex', justifyContent:'center', alignItems:'center', pt: 5, pb: 5}}>UPLOAD NEW IMAGE</Typography>
            <Divider />
            <Box sx={{display:"flex", justifyContent:"center", alignItems:"center", pt:2}}>
              <Button variant="contained" component="label" color="primary" sx={{margin:"20px auto", backgroundColor:"rgba(71, 132, 212, 0.8)"}}>
              <AddAPhoto /> New image
              <Input type="file" style={{display:"none"}} name="image" onChange={(e) => {handleImageChange(e)}}/>
              </Button>
                  
                <Button variant="contained" type="submit" color="primary" component="span" onClick={(e) => uploadImage(e, image, setStatus, setError, navigate)} sx={{margin:"20px auto", backgroundColor:"rgba(71, 132, 212, 0.8)"}}>
                  Upload Image
                </Button>   
            </Box>
            {/* STATUS TEXT */}
            <Box sx={{display:"flex", alignItems:'center', justifyContent:'center'}}>
            {status && <Alert severity='success'>{status}</Alert>}
            {error && <Alert severity='warning'>{error}</Alert>}
            </Box>
            {/* IMAGE PREVIEW */}
            <Box sx={{display:"flex", flexDirection:'column', alignItems:"center", justifyContent:'center', pt: 1}}>      
              {imageUrl && image && (
                <>
                  <Typography variant="h4" fontFamily='Dosis' sx={{display:'flex', justifyContent:'center', alignItems:'center', pt: 5, pb: 5}}>Image Preview</Typography>
                  <Avatar src={imageUrl} alt={image.name} sx={{width: 200, height: 200}} />
                </>
              )}
            </Box>
          </Grid>

      </Grid>
  );
}
export default ProfilePage;