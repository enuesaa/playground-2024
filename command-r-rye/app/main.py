from mlx_lm import load, generate

def main():
    model, tokenizer = load("mlx-community/c4ai-command-r-v01-4bit")

    prompt = tokenizer.apply_chat_template(
        [{'role': 'user', 'content': 'hello'}],
        tokenize=False,
        add_generation_prompt=True,
    )
    
    response = generate(
        model, 
        tokenizer,
        prompt=prompt,
        verbose=True,
        temp=0.2,
        max_tokens=12800,
    )
    print(response)
