using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PlayerMovment : MonoBehaviour
{
    // Start is called before the first frame update
    public Rigidbody2D playerRG;
    //hold the spped velocity
    public float playerVelocity ;
    public AudioSource wingsnd;
    public AudioSource goversnd;


    public GameObject EndGameCavans;
    void Start()
    {
        
    }

    // Update is called once per frame
    void Update()
    {
        if (Input.GetKeyDown("enter"))
        {
            //print("Fly!");
            wingsnd.Play();
            playerRG.velocity = Vector2.up * playerVelocity;

        }
    }

    public void OnCollisionEnter2D(Collision2D collision)
    {
        print("you hit something.");
        GameOver.EndGame();
        //display EndGame Penal
        EndGameCavans.SetActive(true);
        goversnd.Play();
    }
}
