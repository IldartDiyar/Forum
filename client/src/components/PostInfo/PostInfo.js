import { Link, useLocation } from 'react-router-dom';
import './PostInfo.css'
import { useSelector } from 'react-redux';
import Typography from '@mui/material/Typography';
import { Avatar } from '@mui/material';
import Box from '@mui/material/Box';
function PostInfo(props) {
    const userData = useSelector((state) => state.auth.username);
    const location = useLocation();
    const { pathname } = location;
    const [dateStr, timeStr] = props.date.split(" ");    
    return (
      <>
      <Link
              to={{
                pathname: `/posts/`,
                search: `?id=${props.postid.toString()}`
              }}
              className='link-style'
      >
      <Box className='post-info-container' sx={{height: {xs:'100px', sm:'150px', md:'180px', lg:'200px' }}}>
        {/* Header */}
        <div className='post-info-header'>
            <Avatar src={props.avatar} sx={{ height:'auto', marginTop:'4px', width:'30%' }}/>
          
          <Typography sx={{ fontSize: { xs: '10px', sm: '15px', md: '18px', lg: '22px', fontFamily:'Forum', letterSpacing:'2px', fontWeight:'700' } }} className='post-info-header-author'>{props.author}</Typography>
          
          <Typography sx={{ fontSize: { xs: '9px', sm: '10px', md: '12px', lg: '14px' } }} className='post-info-header-date'>{dateStr}</Typography>
          
        </div>
        
        {/* Body */}
        <div className='post-info-body'>
          <div className='post-info-body-title'>
            <p>{props.title}</p>
          </div>
            <span className='post-info-body-content'>{props.content}</span>
            <p className='post-info-body-time'>{timeStr}</p>
        </div>
      </Box>
      </Link>
    </>
    );
  }
export default PostInfo;