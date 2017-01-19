package main


const (
  Utility ComponentCode = "Utility"
  Logging ComponentCode = "Logging"
  Peer ComponentCode = "Peer"
  
)

const (
  LoggingUnknownError ReasonCode = "LoggingUnknownError"
  
)

const errorMapping string = `
{
"Utility":
  {"UtilityUnknownError" :
    {"en":"An unknown error has occured"}

  }
"Logging":
  {"LoggingUnknownError" :
    {"en":"An unknown error has occured"}

  }
}

