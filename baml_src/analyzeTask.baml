// Defining a data model.
class Task {
  todoDate string | null @description(#"
    Expected date when the task should be done or completed. Format YYYY-MM-DD
  "#)
  isDone bool @description("if the task is done or completed")
  description string
}

// Create a function to extract the resume from a string.
function AnalizeTask(todaysDate: string, resume: string) -> Task {
  client "DeepSeekChat"
  prompt #"
    Today is {{todaysDate}}

    Clean the description:
    - Remove any text related about dates from the description. 
      Example: Go to doctor next monday -> Go to the doctor
    - Correct the grammar of the description, without changing the original meaning.

    Notes
    - By default the task is not done or completed, only true when explicitly said in the description.
    - If there is not a date mentioned in the description, the todo date must be null

    {{ resume }}

    {{ ctx.output_format }}
  "#
}



// Test the function with a sample resume. Open the VSCode playground to run this.
test goToTheMall {
  functions [AnalizeTask]
  args {
    todaysDate "2025-10-04"
    resume #"Go to the mall next frinday"#
  }
}


test byGroceries {
  functions [AnalizeTask]
  args {
    todaysDate "2025-10-04"
    resume #"Buy bananas, "#
  }
}