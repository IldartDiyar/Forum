import React, { useState } from 'react';
import './Input.css'
import { TextField, MenuItem } from '@mui/material';
import { useSelector } from 'react-redux';
import Button from '@mui/material/Button';
import { useNavigate } from "react-router-dom";
import {Grid, Typography} from '@mui/material';
import {Box} from '@mui/material';
import { createPost } from '../../api/Posts/CreatePost';
import {Alert} from '@mui/material';

const InputForm = (props) => {
    const navigate = useNavigate();
    const [text, setText] = useState('');  
    const [category, setCategory] = useState('Alem');
    const [title, setTitle] = useState('');
    const userData = useSelector((state) => state.auth.username);
    const [status, setStatus] = useState('');
    
    //TODO  FOR FUTURE IMAGE UPLOAD

    // const [image, setImage] = useState(null);
    // const fileSizeLimit = 5000000; //5 MB
    // const handleImageChange = (e) => {
    //   const uploadedFile = e.target.files[0];
    //   if (uploadedFile && uploadedFile.size > fileSizeLimit){
    //     //setError(`File size must be less than ${fileSizeLimit} bytes`)
    //     setImage(null);
    //   } else{
    //     //setError('')
    //     const reader = new FileReader();

    //     reader.readAsDataURL(uploadedFile);
    
    //     reader.onload = () => {
    //       const base64String = reader.result;
    //       setImage(base64String);
    //     };
    //   }
    // };
    const handleTextChange = event => {
      const inputText = event.target.value;
      if (inputText.length <= 3000) {
        setText(inputText);
      }
    };
    const handleTitleChange = event => {
      const inputText = event.target.value;
      if (inputText.length <= 25) {
        setTitle(inputText);
      }
    };

    const handleStatus = () => {
      if (status){
          return (  
              <Alert severity="error">
                  {status}
              </Alert>
          )
      }
  }

    return (
      <Grid 
      container
      direction="row"
      justifyContent="center"
      alignItems="center"
      >
      <Grid item xs={12}>    
        <Typography variant='h4' sx={{textAlign:'center', padding:'10px'}}>Create Post</Typography>
      </Grid>  
      
      <Grid item xs={12}>  
      <form className='input-form' onSubmit={(e) => createPost(e, text, category, title, userData, setStatus, navigate)}>
        <div className='text-field-group'>
        
          <TextField required={true} id="title-form" className='mu-textfield' label="Title" name="title" variant="outlined" sx={{width:"100%"}} value={title} onChange={(e) => handleTitleChange(e)} autoFocus={true}/>
          
          <TextField
            id="outlined-select-currency"
            select
            name='category'
            defaultValue="Alem"
            sx={{ backgroundColor: 'aliceblue', width: "100%" }}
            onChange={(e) => setCategory(e.target.value)}
          >
            {props.categories.map((option) => (
              <MenuItem key={option.categoryId} value={option.categoryName}>
                {option.categoryName}
              </MenuItem>
            ))}
          </TextField>
        </div> 
        
        <div className='textarea-group'>
          <textarea className='textarea-group-input' name='content' required={true} value={text} onChange={(e) => handleTextChange(e)}/>
        </div>
        
        <Typography component="h1" variant="h5" sx={{color:'black'}}>
          {handleStatus()}
        </Typography>
        <Box sx={{display:'flex', alignItems:'center'}}>
          <Button 
            variant='contained' 
            type="submit" 
            size="large"
            sx={{
              margin:"20px auto",
              backgroundColor:"rgba(71, 132, 212, 0.8)"
            }}>
              SUBMIT
            </Button>
        </Box>
        {/* FOR THE FUTURE POSSIBILITY OF AN IMAGE UPLOAD */}
          {/* <Input type="file" name="image" onChange={handleImageChange} sx={{backgroundColor:'white', mx:'auto'}}/> */}
          {/* <Button variant="contained" component="label" color="primary" sx={{margin:"20px auto", backgroundColor:"rgba(71, 132, 212, 0.8)"}}> */}
          {/* {" "} */}
          {/* <AddAPhoto /> Upload a file */}
          {/* <input type="file" hidden /> */}
          {/* </Button> */}
      </form>
      </Grid>
    </Grid>
    );
  }
export default InputForm;