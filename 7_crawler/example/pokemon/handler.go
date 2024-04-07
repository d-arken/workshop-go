package pokemon

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"sync"
)

type Handler struct {
	Repo *Repository
}

func (h *Handler) PopulatePokemonsList(c *gin.Context) {
	pokemons := GetPokemons()

	err := h.Repo.BatchCreate(pokemons, 100)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) PopulatePokemons(c *gin.Context) {
	var pokes []Model
	err := h.Repo.GetAll(&pokes)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	pipe := make(chan *PokeAPIRes, len(pokes))

	wg := &sync.WaitGroup{}
	wg.Add(len(pokes))

	go func() {
		wg.Wait()
		close(pipe)
	}()

	batchSize := 100
	batch := make([]APIResponse, 0, batchSize)

	for _, pokemon := range pokes {
		go func(poke Model) {
			defer wg.Done()
			p := GetPokemon(poke.Url)
			pipe <- p

			// Se quiser fazer algo em batches ao decorrer de um processamento
			if len(batch) == batchSize {
				batch = make([]APIResponse, 0, batchSize)
				// SEND BATCH
			} else {
				batch = append(batch, *p.Pokemon)
			}
		}(pokemon)
	}

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-pipe; ok {
			c.SSEvent("message", msg.Pokemon)
			return true
		}
		return false
	})
}
