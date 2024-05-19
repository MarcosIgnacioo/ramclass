import { UseQueryResult } from '@tanstack/react-query'
import React from 'react'
import Loading from './Loading'
import Error from './Error'
import getUser, { storeInLocal } from '../functions/store'
import Empty from './Empty'
import UserData from '../classes/UserData'
import { useUser } from './UserContext'
import Student from '../classes/Student'

interface Props {
 fetchedData: UseQueryResult<any, Error>;
 cache: Object[];
 nameSpace: string;
 Container: React.ComponentType<any>;
 returnArray: boolean
}

const State = ({ fetchedData, cache, nameSpace, Container, isFiltered }) => {

 let userLocal: UserData | null
 userLocal = (useUser().username == "") ? getUser() : useUser()

 if (!userLocal) return (<h1 className='alert'>No has iniciado sesión</h1>)

 if (fetchedData.isLoading && Container) {
  return <Loading />;
 }

 if (fetchedData.isError && Container) {
  return <Error />;
 }

 if (fetchedData.isSuccess) {
  if (fetchedData.data[nameSpace] === undefined) {
   storeInLocal(fetchedData.data, nameSpace);
   if (!isFiltered || cache === undefined || cache === null) cache = [fetchedData.data];
  } else {
   storeInLocal(fetchedData.data[nameSpace], nameSpace);
   if (!isFiltered || cache === undefined || cache === null) cache = fetchedData.data[nameSpace];
  }
 }

 if (!cache) return null;

 if (!Container) return cache

 if (cache.length === 0) return <Empty />

 if (cache.length === 1) return <Container {...cache[0]} />

 const responseData = cache.map((data, index) => (
  <Container key={index} {...data} />
 ));

 return <>{responseData}</>;
};

export default State;
