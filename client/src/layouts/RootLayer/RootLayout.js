import { Outlet } from "react-router-dom"
import Header from "../../components/Header/Header";
import './RootLayout.css'
import { Grid } from "@mui/material";

export default function RootLayout() {
    
  
  return (
    <div className="main-layout">
        <Header/>
          <Grid container className="body-container">
              <Grid item xs className="background-column-left"/>
              
              <Grid item xs={12} md={8} lg={6} xl={4} className="middle-column" position='relative'>
  
                  <Outlet />
              </Grid>
              
              <Grid item xs className="background-column-right"/>              
          </Grid>
    </div>
  )
}
