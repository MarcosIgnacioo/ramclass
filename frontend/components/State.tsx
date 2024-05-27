import React from 'react'
import Loading from './Loading'
import Error from './Error'
import getUser, { storeInLocal } from '../functions/store'
import Empty from './Empty'
import UserData from '../classes/UserData'
import { useUser } from './UserContext'

const State = ({ fetchedData, cache, nameSpace, Container, isFiltered }) => {

 let userLocal: UserData | null
 userLocal = (useUser().username == "") ? getUser() : useUser()

 if (!userLocal) return ("No has iniciado sesión")

 if (fetchedData.isLoading && Container) {
  return <Loading />;
 }

 if (fetchedData.isError && Container) {
  return <Error error="Ha ocurrido un error inesperado, puedes volver a intentar lo que querías hacer o actualizar la página." />;
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
