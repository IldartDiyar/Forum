import { Link, useNavigate } from "react-router-dom";
import './Header.css';
import { useSelector, useDispatch } from "react-redux";
import { Stack } from "@mui/system";
import DrawerComponent from "../Drawer/Drawer";
import { useMediaQuery } from "@mui/material";
import { useTheme } from "@mui/material";
import { signOutHandler } from "../../api/Authorization/Authorization";

const Header = () => {  
  const handleMouseMovement = (e) => {
    const x = e.pageX - e.target.offsetLeft
      const y = e.pageY - e.target.offsetTop
    
      e.target.style.setProperty('--x', `${ x }px`)
      e.target.style.setProperty('--y', `${ y }px`)
  }
  const userData = useSelector((state) => state.auth.username);
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down("md"));
  /*const isLaptop = useMediaQuery(theme.breakpoints.up("md"))*/
  const isLarge = useMediaQuery(theme.breakpoints.up("lg"))
  return (
      <Stack direction="row" sx={{ width: '100%', height:"10vh" , backgroundColor:"rgb(0, 0, 0)"}} className="main-header-container">  
      {(userData && !isMobile) && 
        <Stack direction="row" gap="2%" justifyContent="flex-start" alignItems="center" className="nav">
          <Link style={{paddingLeft: isLarge? "40px":"20px"}} to="/">
            <button className="button" style={{padding: isLarge? "15px":"10px", fontSize: isLarge? "small":"x-small"}} onMouseMove={handleMouseMovement}>HOME</button>
          </Link>
          <Link to="/my-posts">
            <button className="button" style={{padding: isLarge? "15px":"10px", fontSize: isLarge? "small":"x-small"}} onMouseMove={handleMouseMovement}>MY POSTS</button>
          </Link>
          
          <Link to="/liked-posts">
            <button className="button" style={{padding: isLarge? "15px":"10px", fontSize: isLarge? "small":"x-small"}} onMouseMove={handleMouseMovement}>LIKED POSTS</button>
          </Link>
          <Link to="/createpost">
            <button className="button" style={{padding: isLarge? "15px":"10px", fontSize: isLarge? "small":"x-small"}} onMouseMove={handleMouseMovement}>NEW POST</button>
          </Link>

        </Stack>          
      }
      {(!userData && !isMobile) &&
        <Stack direction="row" gap="2%" justifyContent="flex-start" alignItems="center" className="nav">
          <Link style={{paddingLeft: isLarge? "40px":"20px"}} to="/">
            <button className="button" style={{padding: isLarge? "15px":"10px", fontSize: isLarge? "small":"x-small"}} onMouseMove={handleMouseMovement}>HOME</button>
          </Link>
        </Stack>
      }
      {(userData && isMobile) && <div className="nav"><DrawerComponent isAuth="true"/></div>}
      {(!userData && isMobile) && <div className="nav"><DrawerComponent isAuth="false"/></div>}  
      
      <Stack direction="row" className="title" justifyContent="center" alignItems="center" style={{padding: isLarge? "20px":"10px", fontSize: isLarge? "1.5rem":"1rem"}} onMouseMove={handleMouseMovement}>
        <Link to="/" style={{textDecoration:'none'}}>
          <h1>FORUM</h1> 
        </Link>        
      
      </Stack>

      {!userData && 
        <Stack direction="row" gap="2%" justifyContent="flex-end" alignItems="center" className="authorization">
          <Link to="/signin">
              <button className="button" style={{padding: isLarge? "15px":"10px", fontSize: isLarge? "small":"x-small"}} onMouseMove={handleMouseMovement}>SIGN IN</button>
            </Link>
            <Link style={{paddingRight: isLarge? "40px":"20px"}} to="/signup">
              <button className="button" style={{padding: isLarge? "15px":"10px", fontSize: isLarge? "small":"x-small"}} onMouseMove={handleMouseMovement}>SIGN UP</button>
            </Link>
        </Stack>
      }
      {userData &&
        <Stack direction="row" gap="2%" justifyContent="flex-end" alignItems="center" className="authorization">
          <Link to="/profile">
            <button className="button" style={{padding: isLarge? "15px":"10px", fontSize: isLarge? "small":"x-small"}} onMouseMove={handleMouseMovement}>PROFILE</button>
          </Link>
          <Link style={{paddingRight: isLarge? "40px":"20px"}} to="/signin">
            <button className="button" style={{padding: isLarge? "10px":"5px", fontSize: isLarge? "small":"x-small"}} onMouseMove={handleMouseMovement} onClick={() => signOutHandler(dispatch, navigate)}>LOG OUT</button>
          </Link>
        </Stack>
      }
    </Stack>    
  );
};
export default Header;