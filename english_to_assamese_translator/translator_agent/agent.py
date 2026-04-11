
import os
from google.adk.agents import LlmAgent
from google.adk.runners import Runner
from google.adk.sessions import InMemorySessionService
import asyncio
from dotenv import load_dotenv  # Add this import

# Load environment variables from the .env file into the script
load_dotenv()
root_agent = LlmAgent(
        name="assamese_translator",
        description="A translator agent that translates English text to Assamese.",
        model="gemini-2.5-flash",
        instruction="""You are a translator agent that translates English text to Assamese. 
        When you receive an English text, you should translate it to Assamese and return the translated text.
        If the input text is not in English, you should return an error message saying "Input text is not in English.".""",

)   
async def main():

    #It keeps the track of conversational hostory 
    #And Pass that history to the agent when needed.
    #It allows the agent to have context about previous interactions,
    session_service = InMemorySessionService()
    #It acts a the main Engine
    #It takes care of executing the agent, managing the session,
    #and handling the interactions between the user and the agent.
    #requests are then sent to Google's Gemini servers.
    runner = Runner(agent=root_agent, session_service=session_service)

    # Define the input text
    english_text = "Artificial intelligence is changing the way we interact with technology."
    print(f"Original English: {english_text}\n")

    # Run the agent with the input text
    #The runner will send the input text to the agent,
    # which will process it according to its instructions and
    # return the translated text in Assamese.
    #await is used to wait for the response from the agent asynchronously, 
    #so that user get complete response before moving to next line of code.
    response = await runner.run(input=english_text)
    print(f"Assamese Translation: {response.text}")
if __name__ == "__main__":   
    asyncio.run(main()) 

