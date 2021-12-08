#Your Wishes

We will create the application named YourWishes. YourWishes is the simple application that can add and display your list of wishes

## Usecases

There are 2 usecase in this simple apps

1. Add Wishes. Adding new wishes into your apps with specification

- expecting the message as request input 
- Message length must less or equal to 60 char
- There is some ID for every wish with format WS-XXXX, where X is random digit alphanumeric
- We also record the time when the wishes is created

2. Display Wishes.

- There is no input request
- Return the list of Message with specific time and id

## Solution
- We will use nanoid for the Wishes id
- we use sqlite for database
- We publish the service to restapi

## Step by step

1. Create Directory for application
```
$ mkdir yourwishes
$ cd yourwishes
```

2. Create go mod
```
$ go mod init yourwishes
```

3. Create the usecase
```
$ gogen usecase AddNewWishes
$ gogen usecase DisplayAllWhishes
```

4. Focus on `AddNewWishes` first

4.1 Completing the `AddNewWishes` Inport Request
```
type InportRequest struct {
	Message string
	Now     time.Time
}
```

4.2. Create the whole things in one command
- Repository with func `SaveWishes`, 
- Entity `Wishes`, 
- inject repo into Outport `SaveWishesRepo`, 
- inject and create the boilerplate code into `Interactor`   
```
$ gogen repository SaveWishes Wishes AddNewWishes
```

4.3. Completing the Entity `Wishes`. We are also creating the value object `WishesID` and error `WishesMessageLengthHasExceeded`
```
type Wishes struct {
	ID      vo.WishesID
	Message string
	Now     time.Time
}

type WishesRequest struct {
	RandomID string
	Message  string
	Now      time.Time
}

func NewWishes(req WishesRequest) (*Wishes, error) {

	var obj Wishes
	
	if len(req.Message) > 60 {
		return apperror.WishesMessageLengthHasExceeded
	}

	id, err := vo.NewWishesID(req.RandomID)
	if err != nil {
		return nil, err
	}

	obj.Message = req.Message
	obj.Now = req.Now
	obj.ID = id

	return &obj, nil
}

```

4.3.1. Creating Value Object (with single type string) `WishesID`
```
$ gogen valuestring WishesID
```

4.3.2. Completing the `WishesID` valueobject
```
type WishesID string

func NewWishesID(randomID string) (WishesID, error) {

	var obj = WishesID(fmt.Sprintf("WS-%s", randomID))

	return obj, nil
}

func (r WishesID) String() string {
	return string(r)
}
```

4.3.3. Creating the error `WishesMessageLengthHasExceeded`
```
$ gogen error WishesMessageLengthHasExceeded
```

4.4. Completing the Interactor. We also create the service for generateID, and wrap it to the transaction
```
func (r *addNewWishesInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := service.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		randomID := r.outport.GenerateID(ctx)

		wishesObj, err := entity.NewWishes(entity.WishesRequest{
			RandomID: randomID,
			Message:  req.Message,
			Now:      req.Now,
		})
		if err != nil {
			return err
		}

		err = r.outport.SaveWishes(ctx, wishesObj)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
```

4.4.1. Create `GenerateID` func service and inject it into Outport and add it as boilerplate to `AddNewWishes` usecase by put the `//!` marker in the code
```
$ gogen service GenerateID AddNewWishes
```

4.4.2. Add `repository.WithTransactionDB` manually into Outport and wrap it into interactor
```
type Outport interface {
	repository.SaveWishesRepo
	service.GenerateIDService
	repository.WithTransactionDB
}
```

5. Focus on Display All Wishes

5.1. Create the whole things in one command
- Repository with func `FindAllWishes`,
- inject repo into Outport `FindAllWishesRepo`,
- inject and create the boilerplate code into `Interactor`
```
$ gogen repository FindAllWishes Wishes DisplayAllWishes
```

5.2. Adding repository to Outport
```
type Outport interface {
	repository.FindAllWishesRepo
	repository.WithoutTransactionDB
}
```

5.3. Modify The interactor `DisplayAllWishes`
```
func (r *displayAllWishesInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := service.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		wishesObjs, err := r.outport.FindAllWishes(ctx)
		if err != nil {
			return err
		}

		for _, obj := range wishesObjs {
			res.ListOfWishes = append(res.ListOfWishes, Wishes{
				ID:      obj.ID.String(),
				Message: obj.Message,
				Date:    obj.Now,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
```



6. Create a gateway (Outport implementation)
```
$ gogen gateway localdb
```

7. Create a controller with gin.gonic web framework
```
$ gogen controller restapi AddNewWishes
$ gogen controller restapi DisplayAllWishes
```

8. Run `go mod tidy` to download all the required dependency code
```
$ go mod tidy
```

9. Modify Controller for `AddNewWishes`
```

	type request struct {
		Message string `json:"message"`
	}        

	return func(c *gin.Context) {

		traceID := util.GenerateID()

		ctx := log.Context(c.Request.Context(), traceID)

		var jsonReq request
		
		if err := c.BindJSON(&jsonReq); err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(err, traceID))
			return
		}

		var req addnewwishes.InportRequest
		req.Message = jsonReq.Message
		req.Now = time.Now()
		
		...
		...
		...
```

10. Modify Controller for `DisplayAllWishes`. Remember to remove the bindJson part since we don't have the request input 
```

	type Wishes struct {
		Message string `json:"message"`
		Date    string `json:"date"`
		ID      string `json:"id"`
	}

	type response struct {
		AllWishes []Wishes `json:"all_wishes"`
	}
	
        ...
        ...
        ...	
	
		var req displayallwishes.InportRequest
		//if err := c.BindJSON(&req); err != nil {
		//	log.Error(ctx, err.Error())
		//	c.JSON(http.StatusBadRequest, NewErrorResponse(err, traceID))
		//	return
		//}	

        ...
        ...
        ...

		var jsonRes response

		for _, w := range res.ListOfWishes {
			jsonRes.AllWishes = append(jsonRes.AllWishes, Wishes{
				Message: w.Message,
				Date:    w.Date.Format("2006-01-02 15:04:05"),
				ID:      w.ID,
			})
		}
```

11. Modify Router
```
	r.Router.POST("/wishes", r.authorized(), r.addNewWishesHandler(r.AddNewWishesInport))
	r.Router.GET("/wishes", r.authorized(), r.displayAllWishesHandler(r.DisplayAllWishesInport))
```

12. Create a Registry
```
gogen registry myapp restapi
```

12. Run `go mod tidy` again to download all the required dependency code
```
$ go mod tidy
```

13. Finishing the gateway part
```
package localdb

import (
	"context"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"yourwishes/infrastructure/database"

	"yourwishes/domain/entity"
	"yourwishes/infrastructure/log"
)

type gateway struct {
	*database.GormWithTransactionImpl
	*database.GormWithoutTransactionImpl
}

// NewGateway ...
func NewGateway() *gateway {

	db := database.NewSQLiteDefault()

    

	return &gateway{
		GormWithTransactionImpl: database.NewGormWithTransactionImpl(db),
		GormWithoutTransactionImpl: database.NewGormWithoutTransactionImpl(db),
	}
}

func (r *gateway) SaveWishes(ctx context.Context, obj *entity.Wishes) error {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return err
	}

	err = db.Save(obj).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gateway) GenerateID(ctx context.Context) string {
	log.Info(ctx, "called")

	id, err := gonanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", 4)
	if err != nil {
		return "abcd"
	}

	return id
}

func (r *gateway) FindAllWishes(ctx context.Context) ([]*entity.Wishes, error) {
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	objs := make([]*entity.Wishes, 0)

	err = db.Find(&objs).Error
	if err != nil {
		return nil, err
	}

	return objs, nil
}

```
