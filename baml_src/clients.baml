// Learn more about clients at https://docs.boundaryml.com/docs/snippets/clients/overview

client<llm> DeepSeekChat {
  provider "openai-generic"
  options {
    base_url "https://api.deepseek.com"
    model "deepseek-chat"
    api_key env.DEEPSEEK_API_KEY
  }
}

client<llm> TogetherLlama3_3_70bTurbo {
  provider "openai-generic"
  options {
    base_url "https://api.together.xyz/v1"
    model "meta-llama/Llama-3.3-70B-Instruct-Turbo"
    api_key env.TOGETHER_API_KEY
  }
}