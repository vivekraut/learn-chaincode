package main


const (
  //Utility ComponentCode = "Utility"
  Logging ComponentCode = "Logging"
  //Peer ComponentCode = "Peer"
  
)

const (
  LoggingUnknownError ReasonCode = "LoggingUnknownError"
  
)

const errorMapping string = `
{
"Logging":
  {"LoggingUnknownError" :
    {"en":"An unknown error has occured"}

  }
}

