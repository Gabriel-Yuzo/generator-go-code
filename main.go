package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Estrutura de diretórios e arquivos com placeholders para a entidade
var estrutura = map[string][]string{
	"internal/domain/models":     {"{{.Entity | ToLower}}.go"},
	"internal/domain/repository": {"{{.Entity | ToLower}}_repository.go"},
	"internal/repository":        {"{{.Entity | ToLower}}_repository_impl.go"},
	"internal/usecase":           {"{{.Entity | ToLower}}_usecase.go"},
	"internal/server/handler":    {"{{.Entity | ToLower}}_handler.go"},
	"internal/server/routes":     {"{{.Entity | ToLower}}_handler.go"},
}

// Campo da entidade
type Field struct {
	Name string
	Type string
}

func main() {
	entidade := "User"
	fields := []Field{
		{"ID", "int"},
		{"Name", "string"},
		// Adicione outros campos conforme necessário
	}

	err := criarEstrutura(entidade)
	if err != nil {
		fmt.Printf("Erro ao criar estrutura do projeto: %s\n", err)
		return
	}

	gerarArquivos(entidade, fields)

	fmt.Println("Estrutura de projeto criada com sucesso!")
}

// Cria a estrutura de diretórios e arquivos
func criarEstrutura(entidade string) error {
	for dir, files := range estrutura {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("erro ao criar diretório %s: %w", dir, err)
		}

		for _, file := range files {
			filePath, err := processFileName(filepath.Join(dir, file), entidade)
			if err != nil {
				return fmt.Errorf("erro ao processar nome do arquivo %s: %w", file, err)
			}
			f, err := os.Create(filePath)
			if err != nil {
				return fmt.Errorf("erro ao criar arquivo %s: %w", filePath, err)
			}
			f.Close()
		}
	}

	return nil
}

// Gera arquivos com conteúdo baseado na entidade fornecida
func gerarArquivos(entidade string, fields []Field) {
	templates := map[string]string{
		"internal/domain/models/{{.Entity | ToLower}}.go":                "templates/entity_model.tmpl",
		"internal/domain/repository/{{.Entity | ToLower}}_repository.go": "templates/entity_repository.tmpl",
		"internal/repository/{{.Entity | ToLower}}_repository_impl.go":   "templates/entity_repository_impl.tmpl",
		"internal/usecase/{{.Entity | ToLower}}_usecase.go":              "templates/entity_usecase.tmpl",
		"internal/server/handler/{{.Entity | ToLower}}_handler.go":       "templates/entity_handler.tmpl",
		"internal/server/routes/{{.Entity | ToLower}}route.go":           "templates/entity_handler.tmpl",
	}

	for pathTemplate, tmpl := range templates {
		path, err := processFileName(pathTemplate, entidade)
		if err != nil {
			fmt.Printf("Erro ao processar nome do arquivo %s: %s\n", pathTemplate, err)
			continue
		}
		createFileWithTemplate(path, tmpl, entidade, fields)
	}
}

// Função auxiliar para processar nomes de arquivos com templates
func processFileName(fileTemplate, entidade string) (string, error) {
	t, err := template.New("fileName").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).Parse(fileTemplate)
	if err != nil {
		return "", fmt.Errorf("erro ao parsear template de nome de arquivo: %w", err)
	}

	var result strings.Builder
	data := struct {
		Entity string
	}{
		Entity: entidade,
	}
	err = t.Execute(&result, data)
	if err != nil {
		return "", fmt.Errorf("erro ao executar template de nome de arquivo: %w", err)
	}

	return result.String(), nil
}

// Função auxiliar para criar arquivos com templates
func createFileWithTemplate(path, tmpl, entidade string, fields []Field) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		fmt.Printf("Erro ao parsear template %s: %s\n", tmpl, err)
		return
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo %s: %s\n", path, err)
		return
	}
	defer f.Close()

	data := struct {
		Entity string
		Fields []Field
	}{
		Entity: entidade,
		Fields: fields,
	}

	err = t.Execute(f, data)
	if err != nil {
		fmt.Printf("Erro ao executar template para arquivo %s: %s\n", path, err)
	}
}
