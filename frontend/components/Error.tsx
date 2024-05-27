import React from 'react'

export default function Error(props) {
 console.log(props)
 const { error } = props
 const [fontSizeP, fontsizeH] = (!props.refreshButton) ? ["1rem", "2rem"] : ["2rem", "3rem"]
 return (
  <main>
   <div className='alert error'>
    <div className='error-header' style={{ fontSize: fontsizeH }}>
     ERROR :(
    </div>
    <p id='error-message' style={{ fontSize: fontSizeP }}>
     {error}
     <div className='refresh-button' >
      <button type="button" onClick={() => { window.location.reload() }} hidden={!props.refreshButton}>
       <svg width="92px" height="92px" viewBox="0 0 24.00 24.00" fill="none" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <path d="M4.06189 13C4.02104 12.6724 4 12.3387 4 12C4 7.58172 7.58172 4 12 4C14.5006 4 16.7332 5.14727 18.2002 6.94416M19.9381 11C19.979 11.3276 20 11.6613 20 12C20 16.4183 16.4183 20 12 20C9.61061 20 7.46589 18.9525 6 17.2916M9 17H6V17.2916M18.2002 4V6.94416M18.2002 6.94416V6.99993L15.2002 7M6 20V17.2916" stroke="#f24968" stroke-width="2.4" stroke-linecap="round" stroke-linejoin="round"></path> </g></svg>
      </button>
     </div>
    </p>
   </div>
  </main>
 )
}

