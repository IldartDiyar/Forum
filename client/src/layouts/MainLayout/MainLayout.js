import { useEffect, useState } from "react"
import { Outlet, NavLink } from "react-router-dom"
import Header from '../../components/Header/Header'
import { useRouteLoaderData } from "react-router-dom";

// import Breadcrumbs from "../components/Breadcrumbs"
import './MainLayout.css'
export default function MainLayout() { 
    return (
    <div className="main-layout">
        <nav>
            <Header/>
        </nav>
        <main>
            <div className="body-container">
                <div className="body-posts-container">
                    <Outlet />
                </div>
            </div>
        </main>
    </div>
  )
}

