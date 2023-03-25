import React, { useState } from "react";

import { Link } from "react-router-dom";
import { Drawer, List, ListItem, ListItemText,Divider,IconButton } from "@mui/material";
import MenuIcon from '@mui/icons-material/Menu';
import './Drawer.css'
import { useSelector } from "react-redux";

function DrawerComponent(props) {
  const userData = useSelector((state) => state.auth.username);

  const [openDrawer, setOpenDrawer] = useState(false);
  return (
    <>
      {userData && (
        <>
        <Drawer
        open={openDrawer}
        onClose={() => setOpenDrawer(false)}
        PaperProps={{
          elevation: 8,
          sx: {
            width: "40%",
            height: "100%",
            borderRadius: "0 20px 20px 0",
            backgroundColor: "rgba(84, 107, 156, 0.9)"
          }
        }}
      >
        <List>
        <ListItem onClick={() => setOpenDrawer(false)}>
            <ListItemText>
              <Link to="/" className="drawer-link">Home</Link>
            </ListItemText>
          </ListItem>
          <Divider/>
          <ListItem onClick={() => setOpenDrawer(false)}>
            <ListItemText>
              <Link to="/my-posts" className="drawer-link">My Posts</Link>
            </ListItemText>
          </ListItem>
          <Divider/>
          <ListItem onClick={() => setOpenDrawer(false)}>
            <ListItemText>
              <Link to="/liked-posts" className="drawer-link">Liked Posts</Link>
            </ListItemText>
          </ListItem>
          <Divider/>
          <ListItem onClick={() => setOpenDrawer(false)}>
            <ListItemText>
              <Link to="/createpost" className="drawer-link">New Post</Link>
            </ListItemText>
          </ListItem>
          <Divider/>
        </List>
      </Drawer>
      <IconButton onClick={() => setOpenDrawer(!openDrawer)} className="drawer-icon">
        <MenuIcon />
      </IconButton>
    </>
      )}
    
    {!userData && (
        <>
        <Drawer
        open={openDrawer}
        onClose={() => setOpenDrawer(false)}
        PaperProps={{
          elevation: 8,
          sx: {
            width: "40%",
            height: "100%",
            borderRadius: "0 20px 20px 0",
            backgroundColor: "rgba(84, 107, 156, 0.9)"
          }
        }}
      >
        <List>
        <ListItem onClick={() => setOpenDrawer(false)}>
            <ListItemText>
              <Link to="/" className="drawer-link">Home</Link>
            </ListItemText>
          </ListItem>
        </List>
      </Drawer>
      <IconButton onClick={() => setOpenDrawer(!openDrawer)} className="drawer-icon">
        <MenuIcon />
      </IconButton>
    </>
      )}
    
    </>
      
  );
}
export default DrawerComponent;