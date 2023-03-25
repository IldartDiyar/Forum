import React from 'react'
import InputForm from '../../components/Input-Form/Input-Form'
import { useSelector } from 'react-redux'
import { Navigate } from 'react-router'
import { useLoaderData } from 'react-router'

const CreatePost = () => {
  //const userData = useSelector(state => state.auth.username)
  const {categories} = useLoaderData();
  return <InputForm categories={categories}/>
  
}
export default CreatePost;
