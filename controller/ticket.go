package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/resto/entity"
	"github.com/ptdrpg/resto/lib"
)

type ItemCount struct {
	Items entity.Item `json:"item"`
	Count int         `json:"count"`
}

type TicketRes struct {
	ID       uint            `json:"id"`
	Customer entity.Customer `json:"customer"`
	Items    []ItemCount     `json:"items"`
	Delivery bool            `json:"delivery"`
	Date     string          `json:"date"`
	Updated_at string      `json:"updated_at"`
	Total    int             `json:"total"`
}

func (c *Controller) FindAllTicket(ctx *gin.Context) {
	tickets, err := c.R.FindAllTicket()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var allTickets []TicketRes
	for i := 0; i < len(tickets); i++ {
		var ticketTemp TicketRes
		customers, customErr := c.R.FindUserById(tickets[i].CustomerID)
		if customErr != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": customErr.Error(),
			})
			return
		}
		ticketTemp.ID = tickets[i].ID
		ticketTemp.Customer = customers
		ticketTemp.Delivery = tickets[i].Delivery
		ticketTemp.Total = tickets[i].Total
		findCount, findCountErr := c.R.GetAllItemCount()
		if findCountErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": findCountErr.Error(),
			})
			return
		}
		var allItems []ItemCount
		var myCount []entity.ItemCount
		for j := 0; j < len(findCount); j++ {
			if findCount[j].TicketID == uint(ticketTemp.ID) {
				myCount = append(myCount, findCount[j])
			}
		}
		for j := 0; j < len(myCount); j++ {
			var item ItemCount
			findItem, findItemErr := c.R.FindItemById(int(myCount[j].ItemID))
			if findItemErr != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": findItemErr.Error(),
				})
				break
			}
			item.Items = findItem
			item.Count = myCount[j].Count
			allItems = append(allItems, item)
		}
		ticketTemp.Items = allItems
		allTickets = append(allTickets, ticketTemp)
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": allTickets,
	})
}

func (c *Controller) FindTicketById(ctx *gin.Context) {
	ticketID := ctx.Param("id")
	id, errConv := strconv.Atoi(ticketID)
	if errConv != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errConv.Error(),
		})
		return
	}
	findTicket, err := c.R.FindTicketById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	findCount, findCountErr := c.R.GetAllItemCount()
	if findCountErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": findCountErr.Error(),
		})
		return
	}

	var allItems []ItemCount
	var myCount []entity.ItemCount
	for i := 0; i < len(findCount); i++ {
		if findCount[i].TicketID == uint(id) {
			myCount = append(myCount, findCount[i])
		}
	}

	for i := 0; i < len(myCount); i++ {
		var item ItemCount
		findItem, findItemErr := c.R.FindItemById(int(myCount[i].ItemID))
		if findItemErr != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": findItemErr.Error(),
			})
			break
		}
		item.Items = findItem
		item.Count = myCount[i].Count
		allItems = append(allItems, item)
	}

	var myTicket TicketRes
	myTicket.ID = findTicket.ID
	myTicket.Date = findTicket.Date
	myTicket.Updated_at = findTicket.Updated_at
	myTicket.Total = findTicket.Total
	myTicket.Items = allItems

	customer, customErr := c.R.FindUserById(findTicket.CustomerID)
	if customErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": customErr.Error(),
		})
		return
	}
	myTicket.Customer = customer

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": myTicket	,
	})
}

type TicketCreateInput struct {
	CustomerID int     `json:"customer_id"`
	Items      [][]int `json:"items"`
	Delivery   bool    `json:"delivery"`
}

func (c *Controller) CreateTickets(ctx *gin.Context) {
	var input TicketCreateInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	newdate := time.Now()
	ticket := entity.Ticket{
		CustomerID: input.CustomerID,
		Delivery:   input.Delivery,
		Date:       newdate.String(),
		Updated_at: newdate.String(),
	}

	createErr := c.R.CreateTicket(&ticket)
	if createErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": createErr.Error(),
		})
		return
	}

	var itemCounts []entity.ItemCount
	for i := range input.Items {
		var itemC entity.ItemCount
		itemC.ItemID = uint(input.Items[i][0])
		itemC.Count = input.Items[i][1]
		itemC.TicketID = ticket.ID

		c.R.CreateItemCount(&itemC)
		itemCounts = append(itemCounts, itemC)
	}
	ticket.Items = itemCounts

	findCust, findErr := c.R.FindUserById(input.CustomerID)
	if findErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "customer not found",
		})
		return
	}

	var allItems []ItemCount
	for i := 0; i < len(itemCounts); i++ {
		itemId := itemCounts[i].ItemID
		var countItem ItemCount
		findItem, errFindItem := c.R.FindItemById(int(itemId))
		if errFindItem != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": errFindItem.Error(),
			})
			break
		}
		countItem.Items = findItem
		countItem.Count = itemCounts[i].Count
		allItems = append(allItems, countItem)
	}

	var total int
	for i := 0; i < len(allItems); i++ {
		oneItem := allItems[i].Items.Price * allItems[i].Count
		total = total + oneItem
	}
	ticket.Total = total
	updateTotal := c.R.UpdateTicket(&ticket)
	if updateTotal != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": updateTotal.Error(),
		})
		return
	}

	convertedPoint := lib.Convert2Point(total)
	var updatePoint entity.Customer
	updatePoint.ID = findCust.ID
	updatePoint.Address = findCust.Address
	updatePoint.Age = findCust.Age
	updatePoint.Email = findCust.Email
	updatePoint.Gender = findCust.Gender
	updatePoint.Name = findCust.Name
	updatePoint.Phone_number = findCust.Phone_number
	updatePoint.Point = findCust.Point + convertedPoint

	c.R.UpdateUser(&updatePoint)
	var resTicket TicketRes
	resTicket.ID = ticket.ID
	resTicket.Customer = updatePoint
	resTicket.Delivery = ticket.Delivery
	resTicket.Items = allItems
	resTicket.Total = total
	resTicket.Date = ticket.Date
	resTicket.Updated_at = ticket.Updated_at

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"data": resTicket,
	})
}

func (c *Controller) DeleteTickets(ctx *gin.Context) {
	ticketID := ctx.Param("id")
	id, errConv := strconv.Atoi(ticketID)
	if errConv != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	ticket, findTicketErr := c.R.FindTicketById(id)
	if findTicketErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "ticket doesn't exist",
		})
		return
	}

	customer, findCustomerErr := c.R.FindUserById(ticket.CustomerID)
	if findCustomerErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": findCustomerErr.Error(),
		})
		return
	}

	extractPoint := lib.Convert2Point(ticket.Total)
	customer.Point = customer.Point - extractPoint

	c.R.UpdateUser(&customer)

	deleting := c.R.DeleteTicket(id)
	if deleting != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": deleting.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ticket succefuly deleted",
	})
}
