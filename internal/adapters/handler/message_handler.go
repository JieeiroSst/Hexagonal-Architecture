package handler

import (
	"net/http"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/services"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
   svc services.MessengerService
}

func NewMessageHandler(MessengerService services.MessengerService) *MessageHandler {
   return &MessageHandler{
       svc: MessengerService,
   }
}

func (h *MessageHandler) CreateMessage(ctx *gin.Context) {
   var message domain.Message
   if err := ctx.ShouldBindJSON(&message); err != nil {
       ctx.JSON(http.StatusBadRequest, gin.H{
           "Error": err,
       })

       return
   }

   err := h.svc.CreateMessage(message)
   if err != nil {
       ctx.JSON(http.StatusBadRequest, gin.H{
           "error": err,
       })
       return
   }

   ctx.JSON(http.StatusCreated, gin.H{
       "message": "New message created successfully",
   })
}

func (h *MessageHandler) ReadMessage(ctx *gin.Context) {
   id := ctx.Param("id")
   message, err := h.svc.ReadMessage(id)

   if err != nil {
       ctx.JSON(http.StatusBadRequest, gin.H{
           "error": err.Error(),
       })
       return
   }
   ctx.JSON(http.StatusOK, message)
}

func (h *MessageHandler) ReadMessages(ctx *gin.Context) {

   messages, err := h.svc.ReadMessages()

   if err != nil {
       ctx.JSON(http.StatusBadRequest, gin.H{
           "error": err.Error(),
       })
       return
   }
   ctx.JSON(http.StatusOK, messages)
}


func (h *MessageHandler) UpdateMessage(ctx *gin.Context) {
    id := ctx.Param("id")
    var message domain.Message
    if err := ctx.ShouldBindJSON(&message); err != nil {
         ctx.JSON(http.StatusBadRequest, gin.H{
              "error": err,
         })
         return
    }
    
    err := h.svc.UpdateMessage(id, message)
    if err != nil {
         ctx.JSON(http.StatusBadRequest, gin.H{
              "error": err,
         })
         return
    }
    
    ctx.JSON(http.StatusOK, gin.H{
         "message": "Message updated successfully",
    })
    }


func (h *MessageHandler) DeleteMessage(ctx *gin.Context) {
    id := ctx.Param("id")
    err := h.svc.DeleteMessage(id)
    if err != nil {
         ctx.JSON(http.StatusBadRequest, gin.H{
              "error": err,
         })
         return
    }
    
    ctx.JSON(http.StatusOK, gin.H{
         "message": "Message deleted successfully",
    })
    }