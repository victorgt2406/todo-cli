// Defining a data model.
class Task {
  todoDate string | null @description(#"
    Expected date when the task should be done or completed. Format YYYY-MM-DD
  "#)
  isDone bool @description("if the task is done or completed")
  description string @description("the description of the task well formatted")
}

// Create a function to extract the resume from a string.
function AnalizeTask(todaysDate: string, taskDescription: string) -> Task {
  client "TogetherLlama3_3_70bTurbo"
  prompt #"
    {{ _.role("system") }}
    Today is {{todaysDate}}.

    {{ _.role("user")}}
    From the following task extract the following information:

    {{ taskDescription }}

    {{ ctx.output_format }}

    * The description should not include date or completion related information
  "#
}



// Test the function with a sample resume. Open the VSCode playground to run this.
test goToTheMall {
  functions [AnalizeTask]
  args {
    todaysDate "saturday 2025-10-04"
    taskDescription #"Go to the mall next friday"#
  }
}


test byGroceries {
  functions [AnalizeTask]
  args {
    todaysDate "saturday 2025-10-04"
    taskDescription #"Buy bananas, "#
  }
}

test meetingWithJohn {
  functions [AnalizeTask]
  args {
    todaysDate "saturday 2025-10-04"
    taskDescription #"Meeting with John next thursday"#
  }
}